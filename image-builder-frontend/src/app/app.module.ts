import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ConfigPageComponent } from './pages/config-page/config-page.component';
import { ArchitectureInputComponent } from './components/architecture-input/architecture-input.component';

@NgModule({
  declarations: [
    AppComponent,
    ConfigPageComponent,
    ArchitectureInputComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
