// Code generated from ./lolcode1/lolcode/lolcode.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // lolcode

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"strings"
)
type Tree struct {
	charts.BaseConfiguration
}
// BaselolcodeListener is a complete listener for a parse tree produced by lolcodeParser.
type BaselolcodeListener struct {
	Tree Tree
	Root *opts.TreeData
	nodes []*opts.TreeData
	current *opts.TreeData
	//Flags Flags
	//QFlags QFlags
}
func NewBaseListener() *BaselolcodeListener {
	l := BaselolcodeListener{
	}
	return &l
}
var start = false
var _ lolcodeListener = &BaselolcodeListener{}

type decl struct {
	name string
	DataType string
	score []string
}
type func_gr struct {
	name string
	perehody []*func_gr
}
var fancs []func_gr
var pos_all []string
var decl_list []decl
var stek_declr []string
var flag = 0

func print_pos(pos_all_now []string) string{
	work := ""
	for key, v := range pos_all_now{
		if key == 0{
			work = "global"
			continue
		}
		work += "."+v
	}
	return work
}
var pos_decl_init []string
func Print_decl(){
	file, err := os.Create("result/perems.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	for key:=range decl_list{
		file.WriteString(decl_list[key].DataType + "\t" + decl_list[key].name + "\t" + print_pos(decl_list[key].score) + "\n")
	}
}
func search_decl(name string)  *decl{
	f := false
	for key, v:=range decl_list{
		if decl_list[key].name == name{
			for key1 :=range v.score{
				if (key1 >= len(pos_all) && v.score[key1] != "argument") || (v.score[key1] != "argument" && v.score[key1] != pos_all[key1]){
					f = true
					break
				}
			}
			if f == false{
				return &v
			}
		}
	}
	return nil
}
func search_func(name string)*func_gr{
	for key :=range fancs{
		if fancs[key].name == name{
			return &fancs[key]
		}
	}
	w := func_gr{name: name}
	fancs = append(fancs, w)
	return &w
}
func Ret_funcs(){
	file, err := os.Create("result/funcs.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	for key, v:=range fancs{
		file.WriteString(fancs[key].name + "\n")
		for key1:=range v.perehody{
			file.WriteString(v.name + " -> " + v.perehody[key1].name + "\n")
		}
	}
}
var pos_func []string
var stack_fung []string
// VisitTerminal is called when a terminal node is visited.
func (s *BaselolcodeListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Println(node.GetText() , node.GetSymbol().GetTokenType())
	work := opts.TreeData{Name: node.GetSymbol().GetText()}
	s.current.Children = append(s.current.Children, &work)
	if node.GetSymbol().GetTokenType() == 41 && flag == 1 {
		pos_all[(len(pos_all))-1] += node.GetText()
		flag = 0
		search_func(node.GetText())//fmt.Println(print_pos(pos_all))
		pos_func = append(pos_func, node.GetText())
	}
	if flag == 2 && (len(stek_declr) < 2 || (len(stek_declr) == 2 && node.GetSymbol().GetTokenType() == 7)){
		stek_declr = append(stek_declr, node.GetText())
		pos_decl_init = pos_all
		return
	}
	if len(stek_declr) == 2 && node.GetSymbol().GetTokenType() != 7{
		work1 := decl{name: stek_declr[1], DataType: "NOOB", score: pos_decl_init}
		decl_list = append(decl_list, work1)
		flag = 0
		stek_declr = nil
	}
	if flag == 2 && len(stek_declr) == 3{
		work1 := decl{name: stek_declr[1], score: pos_decl_init}
		if node.GetSymbol().GetTokenType() == 40 {
			if node.GetText() == "WIN" || node.GetText() == "FAIL" {
				work1.DataType = "TROOF"
			} else if strings.Contains(node.GetText(), "\"") == true {
				work1.DataType = "YARN"
			} else if strings.Contains(node.GetText(), ".") == true {
				work1.DataType = "NUMBAR"
			} else {
				work1.DataType = "NUMBR"
			}
		} else if node.GetSymbol().GetTokenType() == 41 {
			w := search_decl(node.GetText())
			if w == nil {
				fmt.Println("compilation error")
				os.Exit(1)
			}
			work1.DataType = w.DataType
		} else {
			work1.DataType = "NOOB"
		}
		decl_list = append(decl_list, work1)
		flag = 0
		stek_declr = nil
	}
	if (node.GetSymbol().GetTokenType() == 17 || node.GetSymbol().GetTokenType() == 18){
		if flag != 5 {
			flag =3
		}
		return
	}
	if flag == 3 && node.GetSymbol().GetTokenType() != 41 {
		fmt.Println("compilation error")
		os.Exit(1)
	} else if flag == 3 {
		work1 := decl{name: node.GetText(), DataType: "NOOB", score: pos_all}
		work1.score = append(work1.score,"argument")
		decl_list = append(decl_list, work1)
		flag = 0
	}
	if flag == 4{
		stack_fung = append(stack_fung, node.GetText())
	}
	if flag == 4 && len(stack_fung) == 2{
		flag = 5
		w := search_func(stack_fung[1])
		w2 := search_func(pos_func[len(pos_func) - 1])
		w2.perehody = append(w2.perehody, w)
		stack_fung = nil
	}
	if flag == 5 && node.GetSymbol().GetTokenType() == 9{
		flag = 0
	}
}

// VisitErrorNode is called when an error node is visited.
func (s *BaselolcodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaselolcodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaselolcodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaselolcodeListener) EnterProgram(ctx *ProgramContext) {
	if !start {
		s.Root = &opts.TreeData{Name: "Program"}
		start = true
	}
	s.current = s.Root
	s.nodes = append(s.nodes, s.Root)
	pos_all = append(pos_all, "global")
	//fmt.Println(print_pos(pos_all))
	w := func_gr{name: "program"}
	fancs = append(fancs, w)
	pos_func = append(pos_func, "program")
}

// ExitProgram is called when production program is exited.
func (s *BaselolcodeListener) ExitProgram(ctx *ProgramContext) {}

// EnterCode_block is called when production code_block is entered.
func (s *BaselolcodeListener) EnterCode_block(ctx *Code_blockContext) {
	node := opts.TreeData{Name: "Code_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCode_block is called when production code_block is exited.
func (s *BaselolcodeListener) ExitCode_block(ctx *Code_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterStatement is called when production statement is entered.
func (s *BaselolcodeListener) EnterStatement(ctx *StatementContext) {
	node := opts.TreeData{Name: "Statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStatement is called when production statement is exited.
func (s *BaselolcodeListener) ExitStatement(ctx *StatementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterLoop is called when production loop is entered.
func (s *BaselolcodeListener) EnterLoop(ctx *LoopContext) {
	node := opts.TreeData{Name: "Loop"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	pos_all = append(pos_all, "Loop ")
	flag = 1
}

// ExitLoop is called when production loop is exited.
func (s *BaselolcodeListener) ExitLoop(ctx *LoopContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	pos_all = pos_all[:len(pos_all)-1]
	fmt.Println(print_pos(pos_all))
}

// EnterDeclaration is called when production declaration is entered.
func (s *BaselolcodeListener) EnterDeclaration(ctx *DeclarationContext) {
	node := opts.TreeData{Name: "Declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	flag = 2
}

// ExitDeclaration is called when production declaration is exited.
func (s *BaselolcodeListener) ExitDeclaration(ctx *DeclarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterPrint_block is called when production print_block is entered.
func (s *BaselolcodeListener) EnterPrint_block(ctx *Print_blockContext) {
	node := opts.TreeData{Name: "Print_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitPrint_block is called when production print_block is exited.
func (s *BaselolcodeListener) ExitPrint_block(ctx *Print_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterIf_block is called when production if_block is entered.
func (s *BaselolcodeListener) EnterIf_block(ctx *If_blockContext) {
	node := opts.TreeData{Name: "If_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	pos_all = append(pos_all, "If")
	fmt.Println(print_pos(pos_all))
}

// ExitIf_block is called when production if_block is exited.
func (s *BaselolcodeListener) ExitIf_block(ctx *If_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	pos_all = pos_all[:len(pos_all)-1]
	fmt.Println(print_pos(pos_all))
}

// EnterElse_if_block is called when production else_if_block is entered.
func (s *BaselolcodeListener) EnterElse_if_block(ctx *Else_if_blockContext) {
	node := opts.TreeData{Name: "Else_if_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	pos_all = append(pos_all, "else if")
	fmt.Println(print_pos(pos_all))
}

// ExitElse_if_block is called when production else_if_block is exited.
func (s *BaselolcodeListener) ExitElse_if_block(ctx *Else_if_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	pos_all = pos_all[:len(pos_all)-1]
	fmt.Println(print_pos(pos_all))
}

// EnterInput_block is called when production input_block is entered.
func (s *BaselolcodeListener) EnterInput_block(ctx *Input_blockContext) {
	node := opts.TreeData{Name: "Input_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInput_block is called when production input_block is exited.
func (s *BaselolcodeListener) ExitInput_block(ctx *Input_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterFunc_decl is called when production func_decl is entered.
func (s *BaselolcodeListener) EnterFunc_decl(ctx *Func_declContext) {
	node := opts.TreeData{Name: "Func_decl"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	pos_all = append(pos_all, "func ")
	flag = 1
}

// ExitFunc_decl is called when production func_decl is exited.
func (s *BaselolcodeListener) ExitFunc_decl(ctx *Func_declContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	pos_all = pos_all[:len(pos_all)-1]
	fmt.Println(print_pos(pos_all))
	pos_func = pos_func[:len(pos_func)-1]
}

// EnterAssignment is called when production assignment is entered.
func (s *BaselolcodeListener) EnterAssignment(ctx *AssignmentContext) {
	node := opts.TreeData{Name: "Assignment"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAssignment is called when production assignment is exited.
func (s *BaselolcodeListener) ExitAssignment(ctx *AssignmentContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterExpression is called when production expression is entered.
func (s *BaselolcodeListener) EnterExpression(ctx *ExpressionContext) {
	node := opts.TreeData{Name: "Expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitExpression is called when production expression is exited.
func (s *BaselolcodeListener) ExitExpression(ctx *ExpressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEquals is called when production equals is entered.
func (s *BaselolcodeListener) EnterEquals(ctx *EqualsContext) {
	node := opts.TreeData{Name: "Equals"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEquals is called when production equals is exited.
func (s *BaselolcodeListener) ExitEquals(ctx *EqualsContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterNot_equals is called when production not_equals is entered.
func (s *BaselolcodeListener) EnterNot_equals(ctx *Not_equalsContext) {
	node := opts.TreeData{Name: "Not_equals"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitNot_equals is called when production not_equals is exited.
func (s *BaselolcodeListener) ExitNot_equals(ctx *Not_equalsContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterBoth is called when production both is entered.
func (s *BaselolcodeListener) EnterBoth(ctx *BothContext) {
	node := opts.TreeData{Name: "Both"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBoth is called when production both is exited.
func (s *BaselolcodeListener) ExitBoth(ctx *BothContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEither is called when production either is entered.
func (s *BaselolcodeListener) EnterEither(ctx *EitherContext) {
	node := opts.TreeData{Name: "Either"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEither is called when production either is exited.
func (s *BaselolcodeListener) ExitEither(ctx *EitherContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterGreater is called when production greater is entered.
func (s *BaselolcodeListener) EnterGreater(ctx *GreaterContext) {
	node := opts.TreeData{Name: "Greater"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitGreater is called when production greater is exited.
func (s *BaselolcodeListener) ExitGreater(ctx *GreaterContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterLess is called when production less is entered.
func (s *BaselolcodeListener) EnterLess(ctx *LessContext) {
	node := opts.TreeData{Name: "Less"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitLess is called when production less is exited.
func (s *BaselolcodeListener) ExitLess(ctx *LessContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterAdd is called when production add is entered.
func (s *BaselolcodeListener) EnterAdd(ctx *AddContext) {
	node := opts.TreeData{Name: "Add"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAdd is called when production add is exited.
func (s *BaselolcodeListener) ExitAdd(ctx *AddContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSub is called when production sub is entered.
func (s *BaselolcodeListener) EnterSub(ctx *SubContext) {
	node := opts.TreeData{Name: "Sub"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSub is called when production sub is exited.
func (s *BaselolcodeListener) ExitSub(ctx *SubContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMul is called when production mul is entered.
func (s *BaselolcodeListener) EnterMul(ctx *MulContext) {
	node := opts.TreeData{Name: "Mul"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)

}
// ExitMul is called when production mul is exited.
func (s *BaselolcodeListener) ExitMul(ctx *MulContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterDiv is called when production div is entered.
func (s *BaselolcodeListener) EnterDiv(ctx *DivContext) {
	node := opts.TreeData{Name: "Div"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDiv is called when production div is exited.
func (s *BaselolcodeListener) ExitDiv(ctx *DivContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMod is called when production mod is entered.
func (s *BaselolcodeListener) EnterMod(ctx *ModContext) {
	node := opts.TreeData{Name: "Mod"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMod is called when production mod is exited.
func (s *BaselolcodeListener) ExitMod(ctx *ModContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterCast is called when production cast is entered.
func (s *BaselolcodeListener) EnterCast(ctx *CastContext) {
	node := opts.TreeData{Name: "Cast"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCast is called when production cast is exited.
func (s *BaselolcodeListener) ExitCast(ctx *CastContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterAll_ is called when production all_ is entered.
func (s *BaselolcodeListener) EnterAll_(ctx *All_Context) {
	node := opts.TreeData{Name: "All_"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAll_ is called when production all_ is exited.
func (s *BaselolcodeListener) ExitAll_(ctx *All_Context) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterAny_ is called when production any_ is entered.
func (s *BaselolcodeListener) EnterAny_(ctx *Any_Context) {
	node := opts.TreeData{Name: "Any_"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAny_ is called when production any_ is exited.
func (s *BaselolcodeListener) ExitAny_(ctx *Any_Context) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterNot_ is called when production not_ is entered.
func (s *BaselolcodeListener) EnterNot_(ctx *Not_Context) {
	node := opts.TreeData{Name: "Not_"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitNot_ is called when production not_ is exited.
func (s *BaselolcodeListener) ExitNot_(ctx *Not_Context) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterFunc_ is called when production func_ is entered.
func (s *BaselolcodeListener) EnterFunc_(ctx *Func_Context) {
	node := opts.TreeData{Name: "Func_"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	flag = 4
}

// ExitFunc_ is called when production func_ is exited.
func (s *BaselolcodeListener) ExitFunc_(ctx *Func_Context) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}
