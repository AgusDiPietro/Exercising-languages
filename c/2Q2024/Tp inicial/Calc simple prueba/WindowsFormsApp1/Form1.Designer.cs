namespace WindowsFormsApp1
{
    partial class frmcalculadora
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.txtop1 = new System.Windows.Forms.TextBox();
            this.txtop2 = new System.Windows.Forms.TextBox();
            this.txtrta = new System.Windows.Forms.TextBox();
            this.Btnsuma = new System.Windows.Forms.Button();
            this.Btnresta = new System.Windows.Forms.Button();
            this.Btnprod = new System.Windows.Forms.Button();
            this.Lblop1 = new System.Windows.Forms.Label();
            this.Lblop2 = new System.Windows.Forms.Label();
            this.Lbirta = new System.Windows.Forms.Label();
            this.Btncoc = new System.Windows.Forms.Button();
            this.Btnsalir = new System.Windows.Forms.Button();
            this.SuspendLayout();
            // 
            // txtop1
            // 
            this.txtop1.Location = new System.Drawing.Point(398, 96);
            this.txtop1.Name = "txtop1";
            this.txtop1.Size = new System.Drawing.Size(75, 20);
            this.txtop1.TabIndex = 0;
            // 
            // txtop2
            // 
            this.txtop2.Location = new System.Drawing.Point(398, 131);
            this.txtop2.Name = "txtop2";
            this.txtop2.Size = new System.Drawing.Size(75, 20);
            this.txtop2.TabIndex = 1;
            // 
            // txtrta
            // 
            this.txtrta.AcceptsReturn = true;
            this.txtrta.Location = new System.Drawing.Point(398, 170);
            this.txtrta.Name = "txtrta";
            this.txtrta.Size = new System.Drawing.Size(75, 20);
            this.txtrta.TabIndex = 2;
            // 
            // Btnsuma
            // 
            this.Btnsuma.Location = new System.Drawing.Point(479, 93);
            this.Btnsuma.Name = "Btnsuma";
            this.Btnsuma.Size = new System.Drawing.Size(75, 23);
            this.Btnsuma.TabIndex = 3;
            this.Btnsuma.Text = "Suma";
            this.Btnsuma.UseVisualStyleBackColor = true;
            this.Btnsuma.Click += new System.EventHandler(this.button1_Click);
            // 
            // Btnresta
            // 
            this.Btnresta.Location = new System.Drawing.Point(479, 129);
            this.Btnresta.Name = "Btnresta";
            this.Btnresta.Size = new System.Drawing.Size(75, 23);
            this.Btnresta.TabIndex = 4;
            this.Btnresta.Text = "Resta";
            this.Btnresta.UseVisualStyleBackColor = true;
            this.Btnresta.Click += new System.EventHandler(this.Btnresta_Click);
            // 
            // Btnprod
            // 
            this.Btnprod.Location = new System.Drawing.Point(479, 168);
            this.Btnprod.Name = "Btnprod";
            this.Btnprod.Size = new System.Drawing.Size(75, 23);
            this.Btnprod.TabIndex = 5;
            this.Btnprod.Text = "Producto";
            this.Btnprod.UseVisualStyleBackColor = true;
            this.Btnprod.Click += new System.EventHandler(this.Btnprod_Click);
            // 
            // Lblop1
            // 
            this.Lblop1.AutoSize = true;
            this.Lblop1.Location = new System.Drawing.Point(325, 99);
            this.Lblop1.Name = "Lblop1";
            this.Lblop1.Size = new System.Drawing.Size(67, 13);
            this.Lblop1.TabIndex = 6;
            this.Lblop1.Text = "1er operador";
            this.Lblop1.Click += new System.EventHandler(this.label1_Click);
            // 
            // Lblop2
            // 
            this.Lblop2.AutoSize = true;
            this.Lblop2.Location = new System.Drawing.Point(325, 138);
            this.Lblop2.Name = "Lblop2";
            this.Lblop2.Size = new System.Drawing.Size(70, 13);
            this.Lblop2.TabIndex = 7;
            this.Lblop2.Text = "2do operador";
            this.Lblop2.Click += new System.EventHandler(this.label2_Click);
            // 
            // Lbirta
            // 
            this.Lbirta.AutoSize = true;
            this.Lbirta.Location = new System.Drawing.Point(337, 175);
            this.Lbirta.Name = "Lbirta";
            this.Lbirta.Size = new System.Drawing.Size(55, 13);
            this.Lbirta.TabIndex = 8;
            this.Lbirta.Text = "Resultado";
            this.Lbirta.Click += new System.EventHandler(this.label3_Click);
            // 
            // Btncoc
            // 
            this.Btncoc.Location = new System.Drawing.Point(479, 197);
            this.Btncoc.Name = "Btncoc";
            this.Btncoc.Size = new System.Drawing.Size(75, 23);
            this.Btncoc.TabIndex = 9;
            this.Btncoc.Text = "Cociente";
            this.Btncoc.UseVisualStyleBackColor = true;
            this.Btncoc.Click += new System.EventHandler(this.Btncoc_Click);
            // 
            // Btnsalir
            // 
            this.Btnsalir.Location = new System.Drawing.Point(398, 212);
            this.Btnsalir.Name = "Btnsalir";
            this.Btnsalir.Size = new System.Drawing.Size(75, 23);
            this.Btnsalir.TabIndex = 10;
            this.Btnsalir.Text = "Salir";
            this.Btnsalir.UseVisualStyleBackColor = true;
            this.Btnsalir.Click += new System.EventHandler(this.button5_Click);
            // 
            // frmcalculadora
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(634, 259);
            this.Controls.Add(this.Btnsalir);
            this.Controls.Add(this.Btncoc);
            this.Controls.Add(this.Lbirta);
            this.Controls.Add(this.Lblop2);
            this.Controls.Add(this.Lblop1);
            this.Controls.Add(this.Btnprod);
            this.Controls.Add(this.Btnresta);
            this.Controls.Add(this.Btnsuma);
            this.Controls.Add(this.txtrta);
            this.Controls.Add(this.txtop2);
            this.Controls.Add(this.txtop1);
            this.Name = "frmcalculadora";
            this.Text = "Calculadora Simple";
            this.Load += new System.EventHandler(this.Form1_Load);
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.TextBox txtop1;
        private System.Windows.Forms.TextBox txtop2;
        private System.Windows.Forms.TextBox txtrta;
        private System.Windows.Forms.Button Btnsuma;
        private System.Windows.Forms.Button Btnresta;
        private System.Windows.Forms.Button Btnprod;
        private System.Windows.Forms.Label Lblop1;
        private System.Windows.Forms.Label Lblop2;
        private System.Windows.Forms.Label Lbirta;
        private System.Windows.Forms.Button Btncoc;
        private System.Windows.Forms.Button Btnsalir;
    }
}

