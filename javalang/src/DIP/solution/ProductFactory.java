package DIP.solution;

public class ProductFactory {
	public static ProductRepostitory create() {
		return new SQLProductRepository();
	}
}
