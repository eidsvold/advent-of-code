import re


def primer_y_ultimo_digitos(texto: str) -> int:
    digitos = re.sub("[^0-9]", "", texto)
    return int(digitos[0] + digitos[-1])


def resumir_digitos(camino_al_input: str) -> int:
    suma = 0

    with open(camino_al_input) as archivo:
        for linea in archivo:
            suma += primer_y_ultimo_digitos(linea)

    return suma
