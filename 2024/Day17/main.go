package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ysanson/AdventOfCode/pkg"
	"github.com/ysanson/AdventOfCode/pkg/execute"
)

type Registers struct {
	a, b, c uint64
}

func newRegister(aVal uint64) Registers {
	return Registers{aVal, 0, 0}
}

type Instructions struct {
	instructions       []int
	operands           []int
	instructionPointer int
	output             []int
}

func (i *Instructions) RestartAnew() Instructions {
	return Instructions{
		instructions:       slices.Clone(i.instructions),
		operands:           slices.Clone(i.operands),
		instructionPointer: 0,
		output:             make([]int, 0, 10),
	}
}

const (
	ADV = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type Tuple struct {
	l, a int
}

func parseProgram(input string) (Registers, Instructions, []int) {
	lines := strings.Split(input, "\n")
	aReg, _ := strconv.ParseUint(strings.Split(lines[0], " ")[2], 10, 64)
	bReg, _ := strconv.ParseUint(strings.Split(lines[1], " ")[2], 10, 64)
	cReg, _ := strconv.ParseUint(strings.Split(lines[2], " ")[2], 10, 64)
	registers := Registers{aReg, bReg, cReg}
	operations := strings.Split(lines[4], " ")[1]
	instr, op := make([]int, 0, len(operations)/2), make([]int, 0, len(operations)/2)
	fullProg := make([]int, 0, len(operations))
	for i, opcode := range strings.Split(operations, ",") {
		if i%2 == 0 {
			instr = append(instr, pkg.MustAtoi(opcode))
		} else {
			op = append(op, pkg.MustAtoi(opcode))
		}
		fullProg = append(fullProg, pkg.MustAtoi(opcode))
	}
	instrs := Instructions{
		instructions:       instr,
		operands:           op,
		instructionPointer: 0,
		output:             make([]int, 0, 10),
	}
	return registers, instrs, fullProg
}

func joinOutput(output []int) string {
	str := make([]string, len(output))
	for i, n := range output {
		str[i] = strconv.Itoa(n)
	}
	return strings.Join(str, ",")
}

func getComboOperandValue(operand int, registers Registers) uint64 {
	switch operand {
	case 0, 1, 2, 3:
		return uint64(operand)
	case 4:
		return registers.a
	case 5:
		return registers.b
	case 6:
		return registers.c
	default:
		panic("Combo operand is invalid value")
	}
}

func performDivision(registers Registers, operand int) uint64 {
	den := pkg.UintPow(2, getComboOperandValue(operand, registers))
	return registers.a / den
}

func processInstruction(registers *Registers, instr *Instructions) {
	pointer := instr.instructionPointer
	operand := instr.operands[pointer]
	if pointer >= len(instr.instructions) {
		return
	}
	switch instr.instructions[instr.instructionPointer] {
	case ADV: //Combo
		registers.a = performDivision(*registers, operand)
	case BXL: //Literal
		registers.b ^= uint64(operand)
	case BST: //Combo
		registers.b = getComboOperandValue(operand, *registers) % 8
	case JNZ: //Literal
		if registers.a != 0 {
			instr.instructionPointer = instr.operands[pointer] - 1
		}
	case BXC: //Unused
		registers.b ^= registers.c
	case OUT: //Combo
		literalValue := getComboOperandValue(operand, *registers)
		instr.output = append(instr.output, int(literalValue%8))
	case BDV: //Combo
		registers.b = performDivision(*registers, operand)
	case CDV: //Conbo
		registers.c = performDivision(*registers, operand)
	}
	instr.instructionPointer++
}

func runProgram(registers Registers, instr Instructions) []int {
	for instr.instructionPointer < len(instr.instructions) {
		processInstruction(&registers, &instr)
	}
	return instr.output
}

func search(program Instructions, original []int) (seed uint64) {
	for itr := len(original) - 1; itr >= 0; itr-- {
		seed <<= 3
		for !slices.Equal(runProgram(newRegister(seed), program), original[itr:]) {
			seed++
		}
	}
	return
}

func run(input string) (part1 interface{}, part2 interface{}) {
	registers, instructions, fullProg := parseProgram(input)
	part1 = joinOutput(runProgram(registers, instructions))
	part2 = search(instructions, fullProg)
	return
}

func main() {
	execute.Run(run, Tests, Puzzle, false)
}
