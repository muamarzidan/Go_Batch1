#include <stdio.h>

int main(void) {
    int m, v;
    float E;
    printf("Enter the mass of the object (kg): ");
    scanf("%d", &m);
    printf("Enter the velocity of the object (m/s): ");
    scanf("%d", &v);
    E = 0.5 * m * v * v;
    printf("The energy of the object is %.2f J", E);
    return 0;
}
