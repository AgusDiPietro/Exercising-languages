using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace WindowsFormsApp1
{
    public partial class frmcalculadora : Form
    {
        public frmcalculadora()
        {
            InitializeComponent();
        }

        private void label3_Click(object sender, EventArgs e)
        {

        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void label1_Click(object sender, EventArgs e)
        {
            
        }

        private void label2_Click(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            txtrta.Text = (Int32.Parse(txtop1.Text) + Int32.Parse(txtop2.Text)).ToString();
        }

        private void button5_Click(object sender, EventArgs e)
        {

        }

        private void Btnresta_Click(object sender, EventArgs e)
        {
            txtrta.Text = Convert.ToString(Convert.ToInt32(txtop1.Text) - Convert.ToInt32(txtop2.Text));
        }

        private void Btnprod_Click(object sender, EventArgs e)
        {
            txtrta.Text = (Int32.Parse(txtop1.Text) * Int32.Parse(txtop2.Text)).ToString();
        }

        private void Btncoc_Click(object sender, EventArgs e)
        {
            txtrta.Text = (Int32.Parse(txtop1.Text) / Int32.Parse(txtop2.Text)).ToString();
            int num = Int32.Parse(txtop1.Text);
            int den = Int32.Parse(txtop2.Text);
            if (den != 0)
            {
                txtrta.Text = (Int32.Parse(txtop1.Text) / Int32.Parse(txtop2.Text)).ToString();
            }


        }
    }
}
