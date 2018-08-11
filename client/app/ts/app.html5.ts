/*
 * Angular Imports
 */
import {
  NgModule
} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';
import {platformBrowserDynamic} from '@angular/platform-browser-dynamic';
import {
  RouterModule,
  Routes
} from '@angular/router';
import {LocationStrategy, HashLocationStrategy} from '@angular/common';

/*
 * Components
 */
import {KabanApp} from './app.component';
import {HomeComponent} from './components/HomeComponent';
import {AboutComponent} from './components/AboutComponent';

/*
 * Webpack
 */
require('css/styles.css');

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'about', component: AboutComponent },
];

@NgModule({
  declarations: [
    KabanApp,
    HomeComponent,
    AboutComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes)
  ],
  bootstrap: [ KabanApp ],
  providers: [ ]
})
class KabanAppModule {}

platformBrowserDynamic().bootstrapModule(KabanAppModule)
  .catch((err: any) => console.error(err));
