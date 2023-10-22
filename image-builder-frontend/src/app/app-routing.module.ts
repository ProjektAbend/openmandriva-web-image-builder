import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {ConfigPageComponent} from "./pages/config-page/config-page.component";

const routes: Routes = [
  {
    path: '',
    redirectTo: '/build-image',
    pathMatch: 'full'
  },
  {
    path: 'build-image',
    component: ConfigPageComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
