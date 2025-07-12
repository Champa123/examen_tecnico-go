package reader

import (
	"testing"
)

func TestReadCSV(t *testing.T) {

	records, err := ReadCSV("../../transactions.csv")

	if err != nil {
		t.Errorf("error should be nil, is %s", err)
	}

	expected := generateRecords()

	for i, record := range records {
		if i == 0 {
			if record[0] != expected[i][0] && record[1] != expected[i][1] && record[2] != expected[i][2] {
				t.Errorf("Wanted %s, %s, %s; Got: %s, %s, %s", expected[i][0], expected[i][1], expected[i][2],
					record[0], record[1], record[2])
			}
		} else {
			if record[0] != expected[i][0] {
				t.Errorf("Got: %s, wanted: %s", record[0], expected[i][0])
			}
			if record[1] != expected[i][1] {
				t.Errorf("Got: %s, wanted: %s", record[1], expected[i][1])
			}
			if record[2] != expected[i][2] {
				t.Errorf("Got: %s, wanted: %s", record[2], expected[i][2])
			}

		}

	}
}

func TestReadCSVFailure(t *testing.T) {

	_, err := ReadCSV("noExistingPath")

	if err == nil {
		t.Errorf("error should not be nil, is %s", err)
	}

	expectedError := expectedErrorForNoExistingPath()

	if err.Error() != expectedError {
		t.Errorf("Expected error is different than produced error, Got: %s, Wanted: %s", err.Error(), expectedError)
	}

}

func TestReadCSVBrokenTransactions(t *testing.T) {

	_, err := ReadCSV("../../brokenTransactions.csv")

	if err == nil {
		t.Errorf("error should be nil, is %s", err)
	}

	expectedError := expectedErrorForBrokenTransactionsCSV()

	if err.Error() != expectedError {
		t.Errorf("Expected error is different than produced error, Got: %s, Wanted: %s", err.Error(), expectedError)
	}

}

func generateRecords() [][]string {
	return [][]string{
		{"Id", "Date", "Transaction"},
		{"0", "7/15", "+60.5"},
		{"1", "7/28", "-10.3"},
		{"2", "8/2", "-20.46"},
		{"3", "8/13", "+10"},
	}
}

func expectedErrorForNoExistingPath() string {
	return "open noExistingPath: no such file or directory"
}

func expectedErrorForBrokenTransactionsCSV() string {
	return "record on line 2; parse error on line 4, column 4: extraneous or missing \" in quoted-field"
}
