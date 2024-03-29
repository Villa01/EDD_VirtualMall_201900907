import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CargaTiendasComponent } from './tiendas/pages/carga-tiendas/carga-tiendas.component';
import { CatalogoComponent } from './tiendas/pages/catalogo/catalogo.component';
import { ProductosComponent } from './tiendas/pages/productos/productos.component';
import { CarritoComponent } from './tiendas/pages/carrito/carrito.component';
import { LoginComponent } from './tiendas/pages/login/login.component';
import { EliminarCuentaComponent } from './tiendas/pages/eliminar-cuenta/eliminar-cuenta.component';
import { ReportesComponent } from './tiendas/pages/reportes/reportes.component';



const routes: Routes = [
  {
    path: '',
    component: LoginComponent,
    pathMatch: 'full',
  }, 
  {
    path: 'home',
    component: CargaTiendasComponent,
    pathMatch: 'full',
  }, 
  {
    path: 'catalogo',
    component: CatalogoComponent,
    pathMatch: 'full'
  }, 
  {
    path:'productos',
    component: ProductosComponent,
    pathMatch: 'full'
  }, 
  {
    path: 'carrito',
    component: CarritoComponent,
    pathMatch: 'full'
  }, 
  {
    path: 'eliminarCuenta',
    component: EliminarCuentaComponent,
    pathMatch: 'full'
  },
  {
    path: 'reportes',
    component: ReportesComponent,
    pathMatch: 'full'
  }
]


@NgModule({
  declarations: [],
  imports: [
    RouterModule.forRoot(routes),
  ],
  exports: [
    RouterModule
  ]
})
export class AppRoutingModule { }
