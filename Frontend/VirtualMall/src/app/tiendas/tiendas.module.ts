import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';

import { CargaTiendasComponent } from './pages/carga-tiendas/carga-tiendas.component';
import { SharedModule } from '../shared/shared.module';
import { InputJSONComponent } from './components/input-json/input-json.component';



@NgModule({
  declarations: [
    CargaTiendasComponent,
    InputJSONComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    SharedModule
  ], 
  exports: [
    CargaTiendasComponent,
  ]
})
export class TiendasModule { }
