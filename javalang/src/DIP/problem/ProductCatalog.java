package DIP.problem;

import java.util.List;

public class ProductCatalog {
	public void listAllProducts() {
		SQLProductRepository sqlProductRepository = new SQLProductRepository(); 
	
		List<String> allProductNames = sqlProductRepository.getAllProductName();
	
		// Display product names
	}

}
