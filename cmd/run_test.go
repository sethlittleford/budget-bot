package cmd

import "testing"

func TestRunCmdFail(t *testing.T) {
	// budget-bot run must have exactly 1 month argument
	rootCmd.SetArgs([]string{"run"})
	if err := runCmd.Execute(); err == nil {
		t.Fatalf("Run Command should have failed with argument error, but did not")
	}
	rootCmd.SetArgs([]string{"run January December"})
	if err := runCmd.Execute(); err == nil {
		t.Fatalf("Run Command should have failed with argument error, but did not")
	}

	// month arg must be in the ValidArgs field of cobra.Command
	// i.e. in the format of "January", "january", "Jan", "jan", or "1"
	badMonths := []string{"0", "13", "janu", "Januarys", "dece", "Decembers"}
	for _, badMonth := range badMonths {
		rootCmd.SetArgs([]string{"run", badMonth})
		if err := runCmd.Execute(); err == nil {
			t.Fatalf("Run Command should have failed with argument error, but did not")
		}
	}
}
