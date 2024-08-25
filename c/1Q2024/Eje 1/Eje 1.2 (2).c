
// Funci�n para convertir un digito romano a su equivalente decimal
int valor_romano(char c) {
    switch (c) {
        case 'I': return 1;
        case 'V': return 5;
        case 'X': return 10;
        case 'L': return 50;
        case 'C': return 100;
        case 'D': return 500;
        case 'M': return 1000;
        default: return 0;
    }
}

// Funcion para convertir una cadena en numeracion romana a decimal
int romano_a_decimal(const char *romano) {
    int total = 0;
    while (*romano) {
        int valor_actual = valor_romano(*romano);
        int valor_siguiente = valor_romano(*(romano + 1));

        if (valor_actual < valor_siguiente) {
            total -= valor_actual;
        } else {
            total += valor_actual;
        }
        romano++;
    }
    return total;
}

int main() {
	
    char romano[11];

    printf("Ingrese un numero en numeracion romana (hasta 10 caracteres): ");
    scanf("%10s", romano);

    // Convertir el numero romano a decimal
    int decimal = romano_a_decimal(romano);

    printf("El numero %s en numeracion arabiga es: %d\n", romano, decimal);

    return 0;
}

