package picker

import "testing"

func TestRegistryBuiltinsLookupAndOperators(t *testing.T) {
	reg := NewRegistry()

	for _, name := range builtInOperators() {
		f, ok := reg.Lookup(name)
		if !ok || f == nil {
			t.Fatalf("expected built-in operator %q to exist", name)
		}
		if f.Name() != name {
			t.Fatalf("unexpected factory name: got=%q want=%q", f.Name(), name)
		}
	}

	if _, ok := reg.Lookup("does-not-exist"); ok {
		t.Fatal("expected missing operator lookup to return ok=false")
	}

	ops := reg.(*registry).Operators()
	if len(ops) != len(builtInOperators()) {
		t.Fatalf("unexpected operators length: got=%d want=%d", len(ops), len(builtInOperators()))
	}
	for i := 1; i < len(ops); i++ {
		if ops[i-1] > ops[i] {
			t.Fatalf("operators should be sorted: %#v", ops)
		}
	}
}

func TestRegistryRegisterBranches(t *testing.T) {
	reg := NewRegistry()

	if err := reg.Register(nil); err == nil {
		t.Fatal("expected nil factory register error")
	}
	if err := reg.Register(NewStaticFactory("")); err == nil {
		t.Fatal("expected empty-name register error")
	}

	if err := reg.Register(NewStaticFactory("zip-code")); err != nil {
		t.Fatalf("unexpected register error: %v", err)
	}
	if _, ok := reg.Lookup("zip-code"); !ok {
		t.Fatal("expected registered custom operator to be found")
	}

	if err := reg.Register(NewStaticFactory("zip-code")); err == nil {
		t.Fatal("expected duplicate register error")
	}
}

func TestNewStaticFactoryAndBuiltInOperators(t *testing.T) {
	f := NewStaticFactory("my-op")
	if f.Name() != "my-op" {
		t.Fatalf("unexpected static factory name: %q", f.Name())
	}

	ops := builtInOperators()
	if len(ops) == 0 {
		t.Fatal("expected non-empty built-in operators")
	}
}
