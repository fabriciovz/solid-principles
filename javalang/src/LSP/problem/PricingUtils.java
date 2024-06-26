package LSP.problem;

import java.util.ArrayList;
import java.util.List;

public class PricingUtils {
	
	public static void main(String[] args) {

		Product p1 = new Product();
		Product p2 = new Product();
		Product p3 = new InHouseProduct();
	
		List<Product> productList = new ArrayList<>();
		
		productList.add(p1);
		productList.add(p2);
		productList.add(p3);
		
		for (Product product: productList) {
			
			if(product instanceof InHouseProduct) {
				((InHouseProduct) product).applyExtraDiscount();
			}
			System.out.println(product.getDiscount());
			
		}
	}
	
}
