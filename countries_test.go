package countries

import (
	"testing"
)

func TestRegisterAndLookup(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterLocale("en"); err != nil {
		t.Fatalf("RegisterLocale(en): %v", err)
	}
	if err := RegisterLocale("de"); err != nil {
		t.Fatalf("RegisterLocale(de): %v", err)
	}

	if err := SetDefaultLocale("en"); err != nil {
		t.Fatalf("SetDefaultLocale: %v", err)
	}

	val, ok := GetName("US")
	if !ok || val != "United States" {
		t.Errorf("GetName(US) = %q, %v; want United States, true", val, ok)
	}

	// Test auto-uppercase
	val, ok = GetName("us")
	if !ok || val != "United States" {
		t.Errorf("GetName(us) = %q, %v; want United States, true", val, ok)
	}

	val, ok = GetNameForLocale("de", "US")
	if !ok || val == "" {
		t.Errorf("GetNameForLocale(de, US) = %q, %v; want non-empty", val, ok)
	}

	// Test NO (Norway) — important: must not be false/empty
	val, ok = GetName("NO")
	if !ok || val != "Norway" {
		t.Errorf("GetName(NO) = %q, %v; want Norway, true", val, ok)
	}
}

func TestRegisterAllLocales(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterAllLocales(); err != nil {
		t.Fatalf("RegisterAllLocales: %v", err)
	}

	locales := ListRegisteredLocales()
	if len(locales) < 168 {
		t.Errorf("expected at least 168 locales, got %d", len(locales))
	}
}

func TestListLocales(t *testing.T) {
	locales := ListLocales()
	if len(locales) < 168 {
		t.Errorf("expected at least 168 available locales, got %d", len(locales))
	}
}

func TestGetAllNames(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	if err := RegisterLocale("en"); err != nil {
		t.Fatalf("RegisterLocale: %v", err)
	}

	all, err := GetAllNames("en")
	if err != nil {
		t.Fatalf("GetAllNames: %v", err)
	}
	if len(all) < 250 {
		t.Errorf("expected at least 250 countries, got %d", len(all))
	}
}

func TestUnregisteredLocale(t *testing.T) {
	mu.Lock()
	translations = make(map[string]map[string]string)
	defaultLocale = ""
	mu.Unlock()

	err := SetDefaultLocale("zz")
	if err == nil {
		t.Error("SetDefaultLocale(zz) should fail for unregistered locale")
	}

	_, ok := GetNameForLocale("zz", "US")
	if ok {
		t.Error("GetNameForLocale(zz) should return false for unregistered locale")
	}
}
