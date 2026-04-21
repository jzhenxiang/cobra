// Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces.
// In addition to providing an interface, Cobra simultaneously provides a controller to organize your application code.
package cobra

import (
	"github.com/spf13/cobra"
)

// Command is a re-export of spf13/cobra.Command for convenience.
type Command = cobra.Command

// CompletionOptions is a re-export of spf13/cobra.CompletionOptions.
type CompletionOptions = cobra.CompletionOptions

// PositionalArgs is a re-export of spf13/cobra.PositionalArgs.
type PositionalArgs = cobra.PositionalArgs

// FParseErrWhitelist is a re-export of spf13/cobra.FParseErrWhitelist.
type FParseErrWhitelist = cobra.FParseErrWhitelist

// Group is a re-export of spf13/cobra.Group.
type Group = cobra.Group

// ShellCompDirective is a re-export of spf13/cobra.ShellCompDirective.
type ShellCompDirective = cobra.ShellCompDirective

const (
	// ShellCompDirectiveError indicates an error occurred and completions should be ignored.
	ShellCompDirectiveError = cobra.ShellCompDirectiveError

	// ShellCompDirectiveNoSpace indicates that the shell should not add a space after the completion.
	ShellCompDirectiveNoSpace = cobra.ShellCompDirectiveNoSpace

	// ShellCompDirectiveNoFileComp indicates that the shell should not provide file completion.
	ShellCompDirectiveNoFileComp = cobra.ShellCompDirectiveNoFileComp

	// ShellCompDirectiveFilterFileExt indicates that the shell should limit file completion to
	// specific extensions.
	ShellCompDirectiveFilterFileExt = cobra.ShellCompDirectiveFilterFileExt

	// ShellCompDirectiveFilterDirs indicates that only directory names should be provided in
	// file completion.
	ShellCompDirectiveFilterDirs = cobra.ShellCompDirectiveFilterDirs

	// ShellCompDirectiveDefault indicates to let the shell perform its default behavior after
	// completions have been provided.
	ShellCompDirectiveDefault = cobra.ShellCompDirectiveDefault
)

// NoArgs is a re-export of spf13/cobra.NoArgs.
var NoArgs = cobra.NoArgs

// OnlyValidArgs is a re-export of spf13/cobra.OnlyValidArgs.
var OnlyValidArgs = cobra.OnlyValidArgs

// ArbitraryArgs is a re-export of spf13/cobra.ArbitraryArgs.
var ArbitraryArgs = cobra.ArbitraryArgs

// MinimumNArgs returns an error if there are not at least N args.
var MinimumNArgs = cobra.MinimumNArgs

// MaximumNArgs returns an error if there are more than N args.
var MaximumNArgs = cobra.MaximumNArgs

// ExactArgs returns an error if there are not exactly N args.
var ExactArgs = cobra.ExactArgs

// RangeArgs returns an error if the number of args is not within the expected range.
var RangeArgs = cobra.RangeArgs

// MatchAll allows combining multiple PositionalArgs to work in concert.
var MatchAll = cobra.MatchAll

// EnableCommandSorting controls sorting of the slice of commands, which is turned on by default.
var EnableCommandSorting = cobra.EnableCommandSorting

// MousetrapHelpText enables an information splash screen on Windows
// if the CLI is started from explorer.exe.
var MousetrapHelpText = cobra.MousetrapHelpText

// MousetrapDisplayDuration controls how long the MousetrapHelpText message is displayed on Windows.
var MousetrapDisplayDuration = cobra.MousetrapDisplayDuration
