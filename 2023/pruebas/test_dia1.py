from advenimientodelcodigo.dia1 import parte_dos, parte_uno


def test_parte_uno():
    camino_al_input = "./bienes/dia1/input.txt"
    suma = parte_uno.resumir_digitos(camino_al_input)
    assert suma == 55607


def test_parte_dos():
    camino_al_input = "./bienes/dia1/input.txt"
    suma = parte_dos.resumir_digitos(camino_al_input)
    assert suma == 55291
