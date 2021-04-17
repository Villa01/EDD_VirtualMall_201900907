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
import { CalendarioComponent } from './pages/calendario/calendario.component';
import { LoginComponent } from './pages/login/login.component';
import { FormLoginComponent } from './components/form-login/form-login.component';
import { RegistrarComponent } from './components/registrar/registrar.component';
import { EliminarCuentaComponent } from './pages/eliminar-cuenta/eliminar-cuenta.component';
import { EliminarCuentaFormComponent } from './components/eliminar-cuenta-form/eliminar-cuenta-form.component';
import { ReportesComponent } from './pages/reportes/reportes.component';
import { ReporteComponent } from './components/reporte/reporte.component';




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
    FilaProductosComponent,
    CalendarioComponent,
    LoginComponent,
    FormLoginComponent,
    RegistrarComponent,
    EliminarCuentaComponent,
    EliminarCuentaFormComponent,
    ReportesComponent,
    ReporteComponent
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
