import re
from typing import List

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


def separar_por_digitos(texto: str):
    return re.findall(r"(?=([0-9]|" + "|".join(digitos_str) + r"))", texto)


def primer_y_ultimo_digitos(texto: str) -> int:
    digitos = separar_por_digitos(texto)
    primer, ultimo = digitos[0], digitos[-1]

    if primer in digitos_str:
        primer = str(digitos_str.index(primer))
    if ultimo in digitos_str:
        ultimo = str(digitos_str.index(ultimo))

    return int(primer + ultimo)


def resumir_digitos(camino_al_input: str) -> List[str]:
    suma = 0

    with open(camino_al_input) as archivo:
        for linea in archivo:
            suma += primer_y_ultimo_digitos(linea)

    return suma
