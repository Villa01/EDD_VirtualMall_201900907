package main

func OrdenacionBurbuja(lista []int) []int {
	for i:= 1; i< len(lista); i++ {
		for j:= 0; j< len(lista)-1; j++ {
			if lista[i] > lista[j]{
				aux := lista[i]
				lista[i]= lista[j]
				lista[j] = aux
			}

		}
	}
	return lista
}
