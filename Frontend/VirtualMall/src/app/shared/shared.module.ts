import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { FileuploadComponent } from './fileupload/fileupload.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [
    FileuploadComponent,
    SidebarComponent
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
