package main

import (
	"fmt"
	"sync"
)

// Ya que el nombre de esta clase comienza con letra minuscula, esta es unexported, si es mayuscula pues es exported.
type student struct {
	studentID      int
	studentName    string
	studentHobbies []string
}

// pequena funcion  de suma con un loop
func summa(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

// GoRoutines and Channels

var wg = sync.WaitGroup{}

func main() {

	// A travez de los canales la informacion pasa a las rutinas definidas
	ch := make(chan int)
	wg.Add(2)

	// En el parametro de esta funcion dejo saber que estoy enviando informacion al canal.
	// Esto se hace para dar un poco mas de informacion sobre el comportamiento de las funciones
	go func(ch chan<- int) {
		ch <- 25
		wg.Done()
	}(ch)
	// En el parametro de esta funcion dejo saber que estoy sacando informacion del canal.
	// Esto se hace para dar un poco mas de informacion sobre el comportamiento de las funciones
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	wg.Wait()

	//Cuento hasta 5 de manera sincronica dos veces.
	count(5)
	count(5)

	// Cuento hasta 5 de manera asincronica dos veces. go me tranforma las llamadas en GoRoutines
	go count(5)
	go count(5)

	// Esto lo que hace es esperar input a la consola, lo hago para darle tiempo a que mi funcion asincronica ejecute.
	// Deberia existir una manera explicita para arreglar este problemita.
	fmt.Scanln()

	lista := []int{1, 1, 1, 2}
	resultado := summa(lista)
	fmt.Println(resultado)
	// student object
	nStudent := student{
		studentID:      802166192,
		studentName:    "Jose",
		studentHobbies: []string{"Surf", "Study"},
	}
	//accesing object properties
	fmt.Println(nStudent.studentHobbies)

	a := [5]int{1, 2, 3, 4, 5} // Arreglo de tamano 5, debido a que defino el tamano es un arreglo
	b := a                     // Aqui estoy realizando una copia del arreglo a en b
	fmt.Println("Muestro a y b ", a, b)
	a[1] = 1000
	fmt.Println(" Muestro a y b luego de modificar un valor de a ", a, b)
	c := []int{1, 2, 3, 4, 5} // Esto es un 'slice' ya que no defini el tamano
	d := c                    // Aqui le estoy asignando a d la referencia que apunta a c. Tomar en consideracion a la hora de hacer cambios
	fmt.Println("Aqui muestro c y d ", c, d)
	c[1] = 1000
	fmt.Println("Aqui muestro c y d luego de modificarlo un valor de c ", c, d)

}
