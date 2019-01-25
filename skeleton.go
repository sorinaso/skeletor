package skeletor

type Skeleton struct {
	Name        string
	Model       Model
	Environment Environment
	Operations  []Operation
}

func NewSkeleton(
	name string,
	model Model,
	environment Environment,
) Skeleton {
	return Skeleton{Name: name, Model: model, Environment: environment}
}

func (s *Skeleton) Create() error {
	for _, o := range s.Operations {
		if err := o.Execute(); err != nil {
			return err
		}
	}

	return nil
}
