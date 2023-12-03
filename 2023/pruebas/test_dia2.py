from advenimientodelcodigo.dia2 import parte_uno


def test_parte_uno():
    camino_al_input = "./bienes/dia2/input.txt"
    suma = parte_uno.resumir_partidos_validos(camino_al_input)
    assert suma == 2563
