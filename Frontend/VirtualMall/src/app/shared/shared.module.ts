import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { FileuploadComponent } from './fileupload/fileupload.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { RouterModule } from '@angular/router';
import { HeaderComponent } from './header/header.component';
import { MostrarErrorComponent } from './mostrar-error/mostrar-error.component';



@NgModule({
  declarations: [
    FileuploadComponent,
    SidebarComponent,
    HeaderComponent,
    MostrarErrorComponent
  ],
  exports : [
    FileuploadComponent,
    SidebarComponent,
    MostrarErrorComponent
  ],
  imports: [
    CommonModule,
    RouterModule
  ]
})
export class SharedModule { }
