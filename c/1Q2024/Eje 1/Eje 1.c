
#define NUM_PRODUCTS 5

typedef struct {
    int code;
    float price;
} Product;

int main() {
    Product products[NUM_PRODUCTS];
    float total = 0, average;
     float maxPrice = -1, minPrice = 1e9;   //1e9 para asegurar que cualquier precio ingresado sea menor y (-1) para asegurar que cualquier precio ingresado sea mayor.
    int maxPriceCode, minPriceCode;
    int countAboveAvg = 0, countAbove100 = 0;

    // Carga info
    for (int i = 0; i < NUM_PRODUCTS; i++) {
        printf("Ingrese el codigo numerico del articulo %d: ", i + 1);
        scanf("%d", &products[i].code);
        printf("Ingrese el precio del articulo %d: ", i + 1);
        scanf("%f", &products[i].price);
        total += products[i].price;

        // Precio m�x y m�n
        if (products[i].price > maxPrice) {
            maxPrice = products[i].price;
            maxPriceCode = products[i].code;
        }
        if (products[i].price < minPrice) {
            minPrice = products[i].price;
            minPriceCode = products[i].code;
        }
    }

    average = total / NUM_PRODUCTS;

    // Superior al promedio y a $100
    for (int i = 0; i < NUM_PRODUCTS; i++) {
        if (products[i].price > average) {
            countAboveAvg++;
        }
        if (products[i].price > 100) {
            countAbove100++;
        }
    }

    // Ordenar los productos en orden ascendente
    for (int i = 0; i < NUM_PRODUCTS - 1; i++) {
        for (int j = i + 1; j < NUM_PRODUCTS; j++) {
            if (products[i].price > products[j].price) {
                Product temp = products[i];
                products[i] = products[j];
                products[j] = temp;
            }
        }
    }

    // Imprimir resultados
    printf("a. Precio maximo: %.2f, Nro. de articulo: %d\n", maxPrice, maxPriceCode);
    printf("b. Precio minimo: %.2f, Nro. de articulo: %d\n", minPrice, minPriceCode);
    printf("c. Cantidad de articulos con precio superior al precio promedio (%.2f): %d\n", average, countAboveAvg);
    printf("d. Cantidad de articulos con precio superior a $100: %d\n", countAbove100);
    printf("e. Articulos ordenados por precio en orden ascendente:\n");
    for (int i = 0; i < NUM_PRODUCTS; i++) {
        printf("Codigo: %d, Precio: %.2f\n", products[i].code, products[i].price);
    }

    return 0;
}

