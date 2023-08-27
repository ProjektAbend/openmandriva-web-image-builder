import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ConfigPageComponent } from './pages/config-page/config-page.component';
import { ArchitectureInputComponent } from './components/architecture-input/architecture-input.component';
import {FormsModule} from "@angular/forms";
import {MatSelectModule} from "@angular/material/select";
import { NoopAnimationsModule } from '@angular/platform-browser/animations';

@NgModule({
  declarations: [
    AppComponent,
    ConfigPageComponent,
    ArchitectureInputComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    MatSelectModule,
    NoopAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
