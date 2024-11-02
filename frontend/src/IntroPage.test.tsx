import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom'
import IntroPage from './IntroPage';

test('renders CodeBrewery title', () => {
   render(
   <BrowserRouter>
   <IntroPage />
   </BrowserRouter>);
   const titleElement = screen.getByText(/CodeBrewery/i);
   expect(titleElement).toBeInTheDocument();
});
