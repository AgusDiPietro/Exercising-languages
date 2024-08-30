using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tp_Inicial_1._6
{
    internal class Program
    {
        static void Main(string[] args)
        {
            int[] PRECIO = new int[7] { 10, 20, 30, 4, 5, 6, 7 };
            int MONTOTOTAL = 0, monto;
            int tkmax = 0, montomax = 0, A = 0, cat, moto, motomax, i;
            int tk;
            int cant;
            int[] M = new int[5]; //uso M[0] PARA ACUMULAR MONTO TOTAL
            int motomin;
            int prom = 0;
            int pora = 0;
            Console.WriteLine("INGRESE EL NRO DE TICKET");
            tk = Convert.ToInt32(Console.ReadLine());
            while (tk != 0)
            {
                Console.WriteLine("ingrese la categoria de pizza (de 1 a 6)");
                cat = Convert.ToInt32(Console.ReadLine());
                Console.WriteLine("ingrese la cantidad");
                cant = Convert.ToInt32(Console.ReadLine());
                Console.WriteLine("ingrese la moto");
                moto = Convert.ToInt32(Console.ReadLine());
                monto = cant * PRECIO[cat];
                if (tkmax == 0)
                {
                    tkmax = tk;
                    motomax = moto;
                    montomax = monto;
                }
                else
                {
                    if (monto > tkmax)//obtengo monto maximo
                    {
                        tkmax = tk;
                        motomax = moto;
                        montomax = monto;
                    }
                }
                M[moto]++; //acumulo cantidad de viajes por moto
                MONTOTOTAL = MONTOTOTAL + monto; //acumulo monto
                M[0]++; //ACUMULO CANTIDAD DE VIAJES TOTALES
                if (cat == 1)
                {
                    A++;
                }
                Console.WriteLine("INGRESE EL NRO DE TICKET");
                tk = Convert.ToInt32(Console.ReadLine()); //VUELVO a pedir ingreso de ticket paraque ingrese el proximo
                     }//fin ingreso
            prom = MONTOTOTAL / M[0]; //calculo el promedio
            pora = A / M[0];
            motomin = 1; // declaro la moto minima
            for (i = 1; i <= 4; i++)//calculo moto menos viajes
            {
                if (M[i] < M[motomin])
                {
                    motomin = i;
                }
            } //fin calculo moto menos viajes
              // comienzo a mostrar
            Console.WriteLine("el ticket {0} de valor: {1} es el de mayor importe", tkmax,
            montomax);
            Console.WriteLine("la moto {0} es la que hizo menos viajes", motomin);
            Console.WriteLine("el valor promedio de los viajes es {0} ", prom);
            Console.WriteLine("el porcentaje de pedidos de pizzas categoria A es {0}",
            pora);
            Console.WriteLine("la facturacion total del negocio es {0}", MONTOTOTAL);
            Console.ReadKey();
        }
    }
    }
}
