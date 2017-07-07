package main

import "fmt"

type Soldier struct {
}

func (s Soldier) attack() {
	fmt.Println("병사가 적을 공격했다.")
}

func (s Soldier ) detail() {
	fmt.Println("병사는 장군의 명령을 받고 있습니다.")
}
type General struct {

}

func (g General) attack() {
	fmt.Println("장군이 병사에게 공격명령을 내렸다.")
}
func (g General ) detail() {
	fmt.Println("장군은 병사에게 명령을 내리고 있습니다.")
}

type Attacker interface {
	attack()
	detail()
}
func Command(a Attacker) {
	a.attack()
	a.detail()
}