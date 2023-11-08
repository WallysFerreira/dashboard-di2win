import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DestaqueComponent } from './destaque.component';

describe('DestaqueComponent', () => {
  let component: DestaqueComponent;
  let fixture: ComponentFixture<DestaqueComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DestaqueComponent]
    });
    fixture = TestBed.createComponent(DestaqueComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
