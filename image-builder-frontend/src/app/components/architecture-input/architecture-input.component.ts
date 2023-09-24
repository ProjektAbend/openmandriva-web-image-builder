import {Component, EventEmitter, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-architecture-input',
  templateUrl: './architecture-input.component.html',
  styleUrls: ['./architecture-input.component.scss']
})
export class ArchitectureInputComponent implements OnInit {
  options = [
    { value: 'option1', label: 'Option 1' },
    { value: 'option2', label: 'Option 2' },
    { value: 'option3', label: 'Option 3' },
  ];
  @Output() optionChange = new EventEmitter<string>();
  selectedOption: string = this.options[0].value;

  constructor() { }

  ngOnInit(): void {
    this.optionChange.emit(this.selectedOption);
  }

  onOptionChange() {
    this.optionChange.emit(this.selectedOption);
  }

}
