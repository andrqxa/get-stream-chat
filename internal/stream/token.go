package stream

import "time"

// GenerateUserToken generates a Stream Chat user token.
// Token must be created server-side using Stream API secret.
func (c *Client) GenerateUserToken(userID string, ttl time.Duration) (string, error) {
	exp := time.Now().Add(ttl)

	token, err := c.client.CreateToken(userID, exp)
	if err != nil {
		return "", err
	}

	return token, nil
}
