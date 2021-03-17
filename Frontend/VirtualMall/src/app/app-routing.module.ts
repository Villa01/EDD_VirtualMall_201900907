import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CargaTiendasComponent } from './tiendas/pages/carga-tiendas/carga-tiendas.component';
import { CatalogoComponent } from './tiendas/pages/catalogo/catalogo.component';



const routes: Routes = [
  {
    path: '',
    component: CargaTiendasComponent,
    pathMatch: 'full',
  }, 
  {
    path: 'catalogo',
    component: CatalogoComponent,
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
