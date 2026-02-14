package column

import "context"

func (r *Repository) UpdatePosition(ctx context.Context, id int, position float64) error {
	_, err := r.pool.Exec(ctx, "UPDATE columns SET position=$1 WHERE id=$2", position, id)
	if err != nil {
		return err
	}

	return nil
}
