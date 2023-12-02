import re
from typing import List

from ._parte_uno import primer_y_ultimo_como_digitos

digitos_str = [
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
]


def reemplazar_digitos(texto: str) -> int:
    for i, digito in enumerate(digitos_str):
        texto = texto.replace(digito, str(i))

    return texto


def resumir_digitos_real(camino_al_input: str) -> List[str]:
    suma = 0

    with open(camino_al_input) as archivo:
        for linea in archivo:
            suma += primer_y_ultimo_como_digitos(reemplazar_digitos(linea))

    return suma
