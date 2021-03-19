import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { CargaTiendasComponent } from './pages/carga-tiendas/carga-tiendas.component';
import { SharedModule } from '../shared/shared.module';
import { InputJSONComponent } from './components/input-json/input-json.component';
import { TiendaComponent } from './components/tienda/tienda.component';
import { CatalogoComponent } from './pages/catalogo/catalogo.component';
import { ProductosComponent } from './pages/productos/productos.component';
import { ProductoComponent } from './components/producto/producto.component';
import { CarritoComponent } from './pages/carrito/carrito.component';
import { TablaCarritoComponent } from './components/tabla-carrito/tabla-carrito.component';
import { FilaProductosComponent } from './components/fila-productos/fila-productos.component';



@NgModule({
  declarations: [
    CargaTiendasComponent,
    InputJSONComponent,
    TiendaComponent,
    CatalogoComponent,
    ProductosComponent,
    ProductoComponent,
    CarritoComponent,
    TablaCarritoComponent,
    FilaProductosComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    SharedModule
  ], 
  exports: [
    CargaTiendasComponent,
    CatalogoComponent,
  ]
})
export class TiendasModule { }
