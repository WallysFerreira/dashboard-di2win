import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';
import { HomeComponent } from './pages/home/home.component';
import { ChartComponent } from './components/chart/chart.component';
import { MenuComponent } from './components/menu/menu.component';
import { FiltroComponent } from './components/filtro/filtro.component';
import { GraficoComponent } from './pages/grafico/grafico.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    ChartComponent,
    MenuComponent,
    FiltroComponent,
    GraficoComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot([
      {path: '', component: HomeComponent},
      {path: 'grafico', component: GraficoComponent}
    ])
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
