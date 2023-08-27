import { Component, OnInit } from '@angular/core';
import {MatDialog} from "@angular/material/dialog";
import {DialogComponent} from "../../components/dialog/dialog.component";
import {DialogStatus} from "../../models/DialogStatus";

@Component({
  selector: 'app-config-page',
  templateUrl: './config-page.component.html',
  styleUrls: ['./config-page.component.scss']
})
export class ConfigPageComponent implements OnInit {

  constructor(public dialog: MatDialog) {}

  ngOnInit(): void {
  }

  onArchitectureChange(selectedOption: string): void {
    console.log('Selected option:', selectedOption);
  }

  generateImage(): void {
    this.openDialog();
  }

  openDialog(): void {
    this.dialog.open(DialogComponent, {
      width: '350px',
      data: {
        text: "Service has reached maximum capacity. Come back later",
        status: DialogStatus.FAILED
      }
    });
  }
}
