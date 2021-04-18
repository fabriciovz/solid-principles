package DIP.solution;

import java.util.Arrays;
import java.util.List;

public class SQLProductRepository implements ProductRepostitory {
	public List<String> getAllProductName(){
		
		
		return Arrays.asList("soap","toothpaste");
		
	}

}
