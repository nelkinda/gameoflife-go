package gameoflife

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
	"os"
	"reflect"
	"testing"
)

var universe *Universe

func theFollowingUniverse(spec *messages.PickleStepArgument_PickleDocString) error {
	universe = parseSimplifiedLife1_05(spec.Content)
	return nil
}

func theNextGenerationMUSTBe(spec *messages.PickleStepArgument_PickleDocString) error {
	universe = universe.Next()
	expected := parseSimplifiedLife1_05(spec.Content)
	if !reflect.DeepEqual(expected, universe) {
		return fmt.Errorf("Expected: <%s>\nActual: <%s>", expected, universe)
	}
	return nil
}

//noinspection GoUnusedExportedFunction
func InitializeScenario(context *godog.ScenarioContext) {
	context.Step(`^the following universe:$`, theFollowingUniverse)
	context.Step(`^the next generation MUST be:$`, theNextGenerationMUSTBe)
}

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func AssertCoverage() int {
	if testing.CoverMode() != "" {
		if c := testing.Coverage(); c < 1.0 {
			fmt.Printf("Coverage failed at %.1f%%\n", c*100)
			return 1
		}
	}
	return 0
}

func AcceptanceTests() int {
	flag.Parse()
	opts.Paths = flag.Args()
	return godog.TestSuite{
		Name: "gameoflife-go",
		ScenarioInitializer: InitializeScenario,
		Options: &opts,
	}.Run()
}

func TestMain(m *testing.M) {
	// The sequence looks weird but is important for the coverage.
	verifiers := []func() int{
		AcceptanceTests,
		m.Run,
		AssertCoverage,
	}

	status := 0
	for _, verifier := range verifiers {
		if st := verifier(); st > status {
			status = st
		}
	}
	os.Exit(status)
}
