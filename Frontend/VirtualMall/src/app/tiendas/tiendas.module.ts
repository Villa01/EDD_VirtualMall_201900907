import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { CargaTiendasComponent } from './pages/carga-tiendas/carga-tiendas.component';
import { SharedModule } from '../shared/shared.module';
import { InputJSONComponent } from './components/input-json/input-json.component';
import { TiendaComponent } from './components/tienda/tienda.component';
import { CatalogoComponent } from './pages/catalogo/catalogo.component';



@NgModule({
  declarations: [
    CargaTiendasComponent,
    InputJSONComponent,
    TiendaComponent,
    CatalogoComponent
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
