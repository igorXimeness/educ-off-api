func (s SubjectService) CreateSubject(ctx context.Context, subject model.Subject) error {
    // Verifica se já existe uma matéria com o mesmo nome
    existingSubject, err := s.subjectRepo.FindByName(ctx, subject.Name)
    if err != nil {
        return err
    }

    if existingSubject != nil {
        return errors.New("subject with this name already exists")
    }

    // Salva a matéria no repositório
    return s.subjectRepo.Save(ctx, subject)
}
