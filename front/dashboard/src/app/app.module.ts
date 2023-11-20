import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';
import { HomeComponent } from './pages/home/home.component';
import { MainChartComponent } from './components/main-chart/main-chart.component';
import { MenuComponent } from './components/menu/menu.component';
import { FiltroComponent } from './components/filtro/filtro.component';
import { GraficoComponent } from './pages/grafico/grafico.component';
import { DestaqueComponent } from './components/destaque/destaque.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    MainChartComponent,
    MenuComponent,
    FiltroComponent,
    GraficoComponent,
    DestaqueComponent
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
