// Package sonyflake wraps Sonyflake with injectable Config (MachineID from the application).
//
// Design constraints:
//   - MachineID and StartTime come from the caller; no env, viper, or host metadata reads.
//   - No package-level singletons; create a Generator with New and reuse it for the process lifetime.
package sonyflake
