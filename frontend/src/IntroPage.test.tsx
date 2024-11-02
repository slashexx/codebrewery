import { render, screen } from '@testing-library/react';
import IntroPage from './IntroPage';

test('renders CodeBrewery title', () => {
   render(<IntroPage />);
   const titleElement = screen.getByText(/CodeBrewery/i);
   expect(titleElement).toBeInTheDocument();
});
