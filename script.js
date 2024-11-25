let subjects = [];
let modules = [];
let lessons = [];

function renderSubjects() {
    const subjectSelects = document.querySelectorAll('#subject-select, #subject-select-lesson');
    subjectSelects.forEach(select => {
        select.innerHTML = '<option value="">Escolha uma Matéria</option>';
        subjects.forEach(subject => {
            const option = document.createElement('option');
            option.value = subject.name;
            option.textContent = subject.name;
            select.appendChild(option);
        });
    });
}

function renderModules(subjectName) {
    const moduleSelect = document.getElementById('module-select');
    moduleSelect.innerHTML = '<option value="">Escolha um Módulo</option>';

    const filteredModules = modules.filter(module => module.subject === subjectName);
    filteredModules.forEach(module => {
        const option = document.createElement('option');
        option.value = module.name;
        option.textContent = module.name;
        moduleSelect.appendChild(option);
    });
}

document.getElementById('create-subject-form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const subjectName = document.getElementById('subject-name').value;
    
    if (subjects.some(subject => subject.name === subjectName)) {
        alert('Matéria com esse nome já existe!');
        return;
    }

    const newSubject = { name: subjectName };
    subjects.push(newSubject);
    
    renderSubjects();
    
    document.getElementById('subject-name').value = '';
    alert('Matéria criada com sucesso!');
});


document.getElementById('create-module-form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const moduleName = document.getElementById('module-name').value;
    const subjectName = document.getElementById('subject-select').value;
    
  
    if (!subjects.some(subject => subject.name === subjectName)) {
        alert('A matéria selecionada não existe!');
        return;
    }

    if (modules.some(module => module.name === moduleName && module.subject === subjectName)) {
        alert('Módulo com esse nome já existe para essa matéria!');
        return;
    }

    const newModule = { name: moduleName, subject: subjectName };
    modules.push(newModule);
    
    renderModules(subjectName);
    
    document.getElementById('module-name').value = '';
    alert('Módulo criado com sucesso!');
});


document.getElementById('create-lesson-form').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const lessonTitle = document.getElementById('lesson-title').value;
    const lessonContent = document.getElementById('lesson-content').value;
    const subjectName = document.getElementById('subject-select-lesson').value;
    const moduleName = document.getElementById('module-select').value;

    if (!subjects.some(subject => subject.name === subjectName)) {
        alert('A matéria selecionada não existe!');
        return;
    }
    if (!modules.some(module => module.name === moduleName && module.subject === subjectName)) {
        alert('O módulo selecionado não existe para essa matéria!');
        return;
    }

    if (lessons.some(lesson => lesson.title === lessonTitle && lesson.module === moduleName)) {
        alert('Lição com esse título já existe!');
        return;
    }

    const newLesson = { title: lessonTitle, content: lessonContent, subject: subjectName, module: moduleName };
    lessons.push(newLesson);

    document.getElementById('lesson-title').value = '';
    document.getElementById('lesson-content').value = '';
    alert('Lição criada com sucesso!');
});
