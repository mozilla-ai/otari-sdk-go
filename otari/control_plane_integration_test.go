package otari_test

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/mozilla-ai/otari-sdk-go/otari"
	generated "github.com/mozilla-ai/otari-sdk-go/otari/client"
)

// Integration tests for the control-plane surface against a live gateway.
//
// They start a real gateway on SQLite with a master key (control-plane
// endpoints never call an LLM provider, so no provider credentials are needed)
// and drive client.ControlPlane through a full CRUD lifecycle. They skip when
// no gateway command is available; set OTARI_GATEWAY_CMD to run them
// (e.g. "uv run --project ../otari gateway").
//
// Management endpoints authenticate with Authorization: Bearer <master_key>,
// which ControlPlane sends.

const masterKey = "itest-master-key"

func freePort(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return strings.Split(l.Addr().String(), ":")[1]
}

func startGateway(t *testing.T) string {
	t.Helper()
	cmdStr := os.Getenv("OTARI_GATEWAY_CMD")
	if cmdStr == "" {
		cmdStr = "gateway"
	}
	fields := strings.Fields(cmdStr)
	if _, err := exec.LookPath(fields[0]); err != nil {
		t.Skipf("gateway command %q not found; set OTARI_GATEWAY_CMD", fields[0])
	}
	port := freePort(t)
	db := t.TempDir() + "/it.db"
	args := append(append([]string{}, fields[1:]...),
		"serve", "--database-url", "sqlite:///"+db, "--master-key", masterKey,
		"--host", "127.0.0.1", "--port", port, "--auto-migrate", "--log-level", "40")
	cmd := exec.Command(fields[0], args...)
	if err := cmd.Start(); err != nil {
		t.Fatalf("start gateway: %v", err)
	}
	t.Cleanup(func() { _ = cmd.Process.Kill(); _, _ = cmd.Process.Wait() })

	base := "http://127.0.0.1:" + port
	deadline := time.Now().Add(30 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get(base + "/health")
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return base
			}
		}
		time.Sleep(300 * time.Millisecond)
	}
	t.Fatalf("gateway did not become healthy at %s", base)
	return ""
}

func TestControlPlaneLifecycle(t *testing.T) {
	base := startGateway(t)
	client, err := otari.New(otari.WithAPIKey("unused"), otari.WithBaseURL(base))
	if err != nil {
		t.Fatal(err)
	}
	cp := client.ControlPlane(masterKey)
	ctx := context.Background()

	// budgets: create -> get -> update (via ergonomic aliases)
	budget, _, err := cp.Budgets.Create(ctx).
		CreateBudgetRequest(generated.CreateBudgetRequest{MaxBudget: *generated.NewNullableFloat32(ptr(float32(100.0)))}).Execute()
	if err != nil {
		t.Fatalf("create budget: %v", err)
	}
	if budget.BudgetId == "" {
		t.Fatal("expected budget_id on create")
	}
	if _, _, err := cp.Budgets.Get(ctx, budget.BudgetId).Execute(); err != nil {
		t.Fatalf("get budget: %v", err)
	}

	// users: create -> get
	user, _, err := cp.Users.Create(ctx).
		CreateUserRequest(generated.CreateUserRequest{UserId: "itest-user"}).Execute()
	if err != nil {
		t.Fatalf("create user: %v", err)
	}
	if user.UserId != "itest-user" {
		t.Fatalf("expected user_id itest-user, got %q", user.UserId)
	}

	// keys: create returns the secret -> delete
	key, _, err := cp.Keys.Create(ctx).
		CreateKeyRequest(generated.CreateKeyRequest{}).Execute()
	if err != nil {
		t.Fatalf("create key: %v", err)
	}
	if key.Key == "" {
		t.Fatal("expected key secret on create")
	}
	if _, err := cp.Keys.Delete(ctx, key.Id).Execute(); err != nil {
		t.Fatalf("delete key: %v", err)
	}

	// usage: list is readable
	if _, _, err := cp.Usage.List(ctx).Execute(); err != nil {
		t.Fatalf("list usage: %v", err)
	}

	// raw escape hatch: the generated client stays reachable for the full surface.
	if _, _, err := cp.Raw.UsageAPI.ListUsageV1UsageGet(ctx).Execute(); err != nil {
		t.Fatalf("list usage via raw: %v", err)
	}
}

func ptr[T any](v T) *T { return &v }
