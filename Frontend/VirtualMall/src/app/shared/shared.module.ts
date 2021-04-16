import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { FileuploadComponent } from './fileupload/fileupload.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { RouterModule } from '@angular/router';
import { HeaderComponent } from './header/header.component';



@NgModule({
  declarations: [
    FileuploadComponent,
    SidebarComponent,
    HeaderComponent
  ],
  exports : [
    FileuploadComponent,
    SidebarComponent
  ],
  imports: [
    CommonModule,
    RouterModule
  ]
})
export class SharedModule { }
