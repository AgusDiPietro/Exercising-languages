using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TpInicial_1._3
{
    internal class Program
    {
        static void Main(string[] args)
        {
            string cadena;
            //int[]vecval= new int[5];
            int[] vecval = Enumerable.Repeat(0, 5).ToArray();
            int cat, leg, hs, i, sdo, tot;
            tot = 0;
            for (i = 1; i <= 4; i++)
            {
                Console.WriteLine("ingrese el valor de la hs de la categoria {0}", i);
                cadena = Console.ReadLine();
                vecval[i] = Convert.ToInt32(cadena);
            }
            Console.WriteLine("ingrese el nro de legajo");
            cadena = Console.ReadLine();
            leg = Convert.ToInt32(cadena);
            while (leg != 0)
            {
                Console.WriteLine("ingrese el nro de categoria");
                cadena = Console.ReadLine();
                cat = Convert.ToInt32(cadena);
                Console.WriteLine("ingrese el nro de hs trabajadas");
                cadena = Console.ReadLine();
                hs = Convert.ToInt32(cadena);
                sdo = hs * vecval[cat];
                Console.WriteLine("el legajo {0} cobrara {1}", leg, sdo);
                tot = tot + sdo;
                Console.WriteLine("ingrese el nro de legajo");
                cadena = Console.ReadLine();
                leg = Convert.ToInt32(cadena);
            }
            Console.WriteLine("el total a pagar por sueldos es {0}", tot);
            Console.ReadKey();


        }
    }
}
