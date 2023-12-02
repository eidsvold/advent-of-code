from advenimientodelcodigo import dia1


def test_reemplazar_digitos():
    digitos = dia1.reemplazar_digitos("xtwone3four")
    # digitos = dia1.primer_y_ultimo_como_digitos(digitos)
    assert digitos == "x21034"


def test_parte_uno():
    camino_al_input = "./assets/dia1/input.txt"
    suma = dia1.resumir_digitos(camino_al_input)
    assert suma == 55607


def test_parte_dos():
    camino_al_input = "./assets/dia1/input.txt"
    suma = dia1.resumir_digitos_real(camino_al_input)
    assert suma == 54882
