from typing import List, Tuple

limites_de_color = {"red": 12, "green": 13, "blue": 14}


def validar_partido(partido: str) -> Tuple[int, bool]:
    """Determinar si el partido es válido o no."""

    id, juegos = partido.split(":", 1)
    id, juegos = int(id.split(" ")[-1]), juegos.split(";")

    for juego in juegos:
        for cubo in juego.split(","):
            cantidad, color = cubo.strip().split(" ")
            if int(cantidad) > limites_de_color[color]:
                return (id, False)

    return (id, True)


def resumir_partidos_validos(camino_al_input: str) -> int:
    """Resumir la cantidad de partidos válidos."""

    suma = 0

    with open(camino_al_input) as archivo:
        for linea in archivo:
            id, valido = validar_partido(linea)
            if valido:
                suma += id

    return suma
