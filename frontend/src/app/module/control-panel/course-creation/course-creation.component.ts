import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { TeacherService } from '../../../services/teacher.service';
import { Staff } from '../../../model/staff';

@Component({
  selector: 'app-course-creation',
  templateUrl: './course-creation.component.html',
  styleUrls: ['./course-creation.component.css'],
})
export class CourseCreationComponent implements OnInit {
  selectedProgram: string = null;

  programs: any[] = ['Bachelor', 'Master'];

  teachers: any[] = [];
  selectedTeacherAssistant: Staff[];

  creationForm = this.formBuilder.group({
    id: [null],
    name: [undefined, Validators.required],
    program: [this.selectedProgram, Validators.required],
    primary_lector: [undefined as Staff, Validators.required],
    assistant_lector: [undefined as Staff, Validators.required],
    teacher_assistants: [[], Validators.required],
  });

  constructor(private formBuilder: FormBuilder, private teacherService: TeacherService) {}

  ngOnInit(): void {
    this.selectedProgram = this.programs[0];
    this.teacherService.getListOfTeachers().subscribe((value) => {
      this.teachers = value.teachers.map((val) => {
        return {
          name: val.degree + ' ' + val.name + ' ' + val.surname,
        };
      });
    });
  }

  public setSelectedTeacherAssistant(val: any[]): void {
    // restore original order
    if (val !== null) {
      this.selectedTeacherAssistant = this.teachers.filter((teacher) => val.includes(teacher));
    }
  }

  submit(): void {
    this.creationForm.reset();
    this.selectedProgram = this.programs[0];
    this.creationForm.patchValue({
      program: this.selectedProgram,
    });
  }
}
