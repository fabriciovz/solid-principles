package DIP.solution;

import java.util.List;

public class ProductCatalog {
	public void listAllProducts() {
		ProductRepostitory productRepository = ProductFactory.create(); 
	
		List<String> allProductNames = productRepository.getAllProductName();
	
		// Display product names
	}

}
