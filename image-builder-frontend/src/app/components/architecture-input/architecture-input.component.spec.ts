import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ArchitectureInputComponent } from './architecture-input.component';

describe('ArchitectureInputComponent', () => {
  let component: ArchitectureInputComponent;
  let fixture: ComponentFixture<ArchitectureInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ArchitectureInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ArchitectureInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
