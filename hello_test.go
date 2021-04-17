package main

import "testing"

func TestHello(t *testing.T){
	//testing.TB est pour avoir l'interface dest T esting et B enchmark
	assertCorrectMessage:=func(t testing.TB, got, want string){
		//Le helper permet d'avoir le vrai endroit ou le test ne marche pas sinon il indiquera toujours la meme ligne
		t.Helper()
		if got!=want{
			t.Errorf("got %v want %v",got,want)
		}
	}

	t.Run("Saying Hello to people",func(t *testing.T){
		got:=Hello("Fabien","")
		want := "Hello Fabien"

		assertCorrectMessage(t,got,want)
	})

	t.Run("Say hello world when no name is supplied",func(t *testing.T){
		got:=Hello("","")
		want:="Hello world"
		assertCorrectMessage(t,got,want)
	})

	t.Run("Say hello in french",func(t *testing.T){
		got:=Hello("Fabien","fr")
		want:="Bonjour Fabien"
		assertCorrectMessage(t,got,want)
	})

	
	t.Run("Say hello in spanish",func(t *testing.T){
		got:=Hello("Fabien","es")
		want:="Hola Fabien"
		assertCorrectMessage(t,got,want)
	})
}